#!/bin/zsh
set -u

# 1) Siempre correr desde la carpeta del script (evita confusiones de ./app)
SCRIPT_DIR="${0:A:h}"
cd "$SCRIPT_DIR" || { echo "No pude hacer cd a $SCRIPT_DIR"; exit 1; }

APP="./app"
TIMEOUT=1    # segundos
RUNS=5000

TMP="$(mktemp -t stress-test.XXXXXX)" || exit 1

cleanup() {
  rm -f "$APP" "$TMP"
}
trap cleanup EXIT INT TERM

# 2) Compilar una sola vez
go build -o "$APP" . || { echo "Build failed"; exit 1; }

for i in {1..$RUNS}; do
  : > "$TMP"

  # 3) Ejecutar con timeout (sin coreutils). Si cuelga -> exit 124
  perl -e '
    my ($t, $out, @cmd) = @ARGV;

    my $pid = fork();
    die "fork failed\n" unless defined $pid;

    if ($pid == 0) {
      open(STDOUT, ">>", $out) or die "open stdout failed\n";
      open(STDERR, ">>", $out) or die "open stderr failed\n";
      exec @cmd or die "exec failed\n";
    }

    local $SIG{ALRM} = sub { kill 9, $pid; exit 124; };
    alarm($t);

    waitpid($pid, 0);
    alarm(0);

    exit($? >> 8);
  ' "$TIMEOUT" "$TMP" "$APP"

  rc=$?

  if [ $rc -ne 0 ]; then
    if [ $rc -eq 124 ]; then
      echo "CUELGUE (timeout) en iteración $i"
    else
      echo "FALLO (exit=$rc) en iteración $i"
    fi

    echo "---- output de la iteración $i ----"
    tail -n 200 "$TMP"
    echo "-----------------------------------"
    break
  fi

  if (( i % 25 == 0 )); then
    echo "OK hasta $i"
  fi
done