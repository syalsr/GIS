#pragma once

#include <cmath>

namespace Geo
{
    constexpr double PI = 3.1415926535;
    constexpr size_t EARTH_RADIUS = 6371000;
    struct Coordinate
    {
        double latitude;
        double longitude;
    };
    double convert_degrees_to_radians(Coordinate c);
    double distance_between_coordinate(Coordinate lhs, Coordinate rhs);
}
