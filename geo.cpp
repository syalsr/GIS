#include "geo.hpp"

double convert_degrees_to_radians(double c)
{
    return c * Geo::PI / 180;
}

double distance_between_coordinate(Geo::Coordinate lhs, Geo::Coordinate rhs)
{
    Geo::Coordinate lhsr{convert_degrees_to_radians(lhs.latitude),
                         convert_degrees_to_radians(lhs.longitude)
    };
    Geo::Coordinate rhsr{convert_degrees_to_radians(rhs.latitude),
                         convert_degrees_to_radians(rhs.longitude)
    };
    std::acos(std::sin(lhsr.latitude) * std::sin(rhsr.latitude) +
        std::cos(lhsr.latitude) * std::cos(rhsr.latitude) *
        std::cos(std::abs(lhsr.longitude - rhsr.longitude))
        ) * Geo::EARTH_RADIUS;
}