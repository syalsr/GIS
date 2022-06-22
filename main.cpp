#include <iostream>
#include <vector>
#include <set>
#include <string>
#include <map>
#include <unordered_map>

struct Stop
{
    std::set<std::string> bus_names;
};

struct Bus
{
    std::size_t stop_cout{};
    std::size_t unique_stop_count{};
    int length_real_route{};
    double length_geographical_route{};
};

class TransportRouter
{
public:

private:
    std::unordered_map<std::string, Stop> stops_;
    std::unordered_map<std::string, Bus> buses_;
};

int main() {
    std::cout << "Hello, World!" << std::endl;
    std::vector<int> v{1,2,3,4,5,6,};
    return 0;
}
