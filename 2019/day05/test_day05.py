from io import StringIO
from day05 import run_intcode_part2


def test_run_intcode_part2(monkeypatch, capsys):
    intcode_input = StringIO("0\n")
    monkeypatch.setattr("sys.stdin", intcode_input)
    run_intcode_part2("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")
    assert capsys.readouterr().out == "Enter input: 0\n"

    intcode_input = StringIO("7\n")
    monkeypatch.setattr("sys.stdin", intcode_input)
    run_intcode_part2("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")
    assert capsys.readouterr().out == "Enter input: 1\n"

    intcode_input = StringIO("5\n")
    monkeypatch.setattr("sys.stdin", intcode_input)
    run_intcode_part2(
        "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
    )
    assert capsys.readouterr().out == "Enter input: 999\n"

    intcode_input = StringIO("8\n")
    monkeypatch.setattr("sys.stdin", intcode_input)
    run_intcode_part2(
        "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
    )
    assert capsys.readouterr().out == "Enter input: 1000\n"

    intcode_input = StringIO("20\n")
    monkeypatch.setattr("sys.stdin", intcode_input)
    run_intcode_part2(
        "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
    )
    assert capsys.readouterr().out == "Enter input: 1001\n"
