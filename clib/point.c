#include "point.h"
#include "common/common.h"

int sum(int a, int b)
{
    cpoint p = {a, b};
    return p.x + p.y;
}
