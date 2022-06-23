#pragma once

#include <string>
#include <vector>
#include <unordered_map>
#include <json.hpp>

#include "geo.hpp"
#include "TransportDataBase.hpp"

namespace RequestToAdd
{
    struct Stop
    {
        std::string name;
        Geo::Coordinate coordinate;
        std::unordered_map<std::string, int> distance_between_stops;
    };

    struct Bus
    {
        std::string name;
        std::vector<std::string> stops;
    };

    struct Route_settings
    {
        int bus_wait_time{};
        double bus_velocity{};
    };

    Stop ParseStop(const std::unordered_map<std::string, Json::Node>& stop_info);
    Bus ParseBus(const std::unordered_map<std::string, Json::Node>& bus_info);

    std::vector<std::variant<Stop, Bus>>
    ParseBaseRequest(const std::vector<Json::Node>& request);
}

namespace RequestToGet
{
    struct Route
    {
        std::string from;
        std::string to;
        std::size_t id;
    };
    struct Stop
    {
        std::string name;
        std::size_t id;
    };
    struct Bus
    {
        std::string name;
        std::size_t id;
    };

    std::vector<Json::Node> ParseStatRequest(const TransportDataBase& db,
                                             const std::vector<Json::Node>& request
                                             );
}