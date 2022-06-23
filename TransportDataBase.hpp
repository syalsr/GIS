#pragma once
#include <set>
#include <string>
#include <unordered_map>
#include <vector>

#include "request.hpp"

struct Stop
{
    std::string name;
    std::set<std::string> bus_names;
};

struct Bus
{
    std::string name;
    std::size_t stop_cout{};
    std::size_t unique_stop_count{};
    int         length_real_route{};
    double      length_geographical_route{};
};

class TransportDataBase
{
public:
    TransportDataBase(std::vector<std::variant<RequestToAdd::Stop,
                                                RequestToAdd::Bus>>,
                                                RequestToAdd::Route_settings);

private:
    double ComputeRealRouteLength(const std::vector<std::string>& stops);
    double ComputeGeographicalRouteLength();

private:
    std::unordered_map<std::string, Stop> stops_;
    std::unordered_map<std::string, Bus> buses_;
    TransportRouter transport_router_;
};