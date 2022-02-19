# coding=utf-8
import heapq
from pathlib import Path


INPUT_DIR = Path(__file__).resolve().parent.parent


class Graph:
    def __init__(self, risks):
        self.length = len(risks)
        self.width = len(risks[0])
        self.risks = {
            f"{x}-{y}": int(risk)
            for y, line in enumerate(risks)
            for x, risk in enumerate(line)
        }
        self.nodes = self.risks.keys()

    @staticmethod
    def get_neighbours(node_name):
        x, y = map(int, node_name.split("-"))
        return [f"{x-1}-{y}", f"{x+1}-{y}", f"{x}-{y-1}", f"{x}-{y+1}"]

    def get_min_risk(self):
        visited = {}
        start_node = f"0-0"
        stop_node = f"{self.width-1}-{self.length-1}"
        min_heap = [(0, start_node)]
        while True:
            current_risk, current_node = heapq.heappop(min_heap)
            if current_node in visited:
                continue
            if current_node == stop_node:
                return current_risk
            for neighbour in self.get_neighbours(current_node):
                if neighbour not in visited:
                    neighbour_risk = self.risks.get(neighbour)
                    if neighbour_risk:
                        heapq.heappush(
                            min_heap, (current_risk + neighbour_risk, neighbour)
                        )
            visited[current_node] = True


def load_input():
    input_file = INPUT_DIR / f"inputs/15/input.txt"
    with input_file.open("r") as f:
        content = f.read()
    return content.splitlines()


if __name__ == "__main__":
    input_risks = load_input()
    graph = Graph(input_risks)
    solution = graph.get_min_risk()
    print("toto", solution)
