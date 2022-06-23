#include <iostream>
#include <vector>
#include <map>

#include "json.hpp"
#include "TrasportRouter.hpp"
#include "TransportDataBase.hpp"


int main() {
    auto request = Json::Load(std::cin);
    auto request_as_map = request.GetRoot().AsMap();

    TransportDataBase transport_db{};
    return 0;
}
