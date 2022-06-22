#pragma once
#include <set>
#include <string>
#include <unordered_map>
#include <vector>

struct Stop
{
    std::set<std::string> bus_names;
};

struct Bus
{
    std::size_t stop_cout{};
    std::size_t unique_stop_count{};
    int         length_real_route{};
    double      length_geographical_route{};
};

class DataBase
{
public:

private:
    double ComputeRealRouteLength(const std::vector<std::string>& stops);
    double ComputeGeographicalRouteLength

private:
    std::unordered_map<std::string, Stop> stops_;
    std::unordered_map<std::string, Bus> buses_;
};