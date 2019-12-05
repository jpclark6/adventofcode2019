require 'pry'
require 'set'

def instructions
    inst = [[],[]]
    File.readlines('./puzzledata/3day.txt').each_with_index do |line, i|
        wire = line.split(',')
        wire.each do |loc|
            inst[i] << {direction: loc[0], length: loc[1..-1].to_i}
        end
    end
    inst
end

def find_manhattan_dist(point)
    x = point[0]
    y = point[1]
    x.abs + y.abs
end

def first_time_to_location?(loc)
    loc == nil
end

def solution
    wires_coords = [Hash.new(), Hash.new()]
    current_x, current_y, tot_distance = 0, 0, 0
    
    instructions.each_with_index do |bend, wire_number|
        bend.each_with_index do |length, bend_number|
            if bend_number == 0
                current_x, current_y, tot_distance = 0, 0, 0
            end
            dist = length[:length]
            
            case length[:direction]
            when "R"
                (current_x..current_x + dist).each do |coord|
                    if first_time_to_location?(wires_coords[wire_number][[coord, current_y]])
                        wires_coords[wire_number][[coord, current_y]] = tot_distance
                    end
                    tot_distance += 1
                end
                current_x += dist
            when "L"
                current_x.downto(current_x - dist).each do |coord|
                    if first_time_to_location?(wires_coords[wire_number][[coord, current_y]])
                        wires_coords[wire_number][[coord, current_y]] = tot_distance
                    end
                    tot_distance += 1
                end
                current_x -= dist
            when "U"
                (current_y..current_y + dist).each do |coord|
                    if first_time_to_location?(wires_coords[wire_number][[current_x, coord]])
                        wires_coords[wire_number][[current_x, coord]] = tot_distance
                    end
                    tot_distance += 1
                end
                current_y += dist
            when "D"
                current_y.downto(current_y - dist).each do |coord|
                    if first_time_to_location?(wires_coords[wire_number][[current_x, coord]])
                        wires_coords[wire_number][[current_x, coord]] = tot_distance
                    end
                    tot_distance += 1
                end
                current_y -= dist
            end
            tot_distance -= 1
        end
    end

    # part 1
    intersections = wires_coords[0].keys & wires_coords[1].keys
    intersection_distances = intersections.map do |intersecting_point|
        find_manhattan_dist(intersecting_point)
    end
    sorted = intersection_distances.sort
    sorted.shift # remove 0, 0 coordinate
    puts "Answer to part 1:", sorted.first

    # part 2
    steps_to_intersection = []
    steps_to_intersection = intersections.map do |coord|
        wires_coords[0][coord] + wires_coords[1][coord] 
    end
    sorted = steps_to_intersection.sort
    puts "Answer to part 2:", sorted[1]
end

solution