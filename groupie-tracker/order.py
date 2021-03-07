def search_unordered_list(value, elements: list):
    for idx, element in enumerate(elements):
        if element == value :
            return idx   
    return -1


    