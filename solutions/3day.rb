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

def part_1
    all_coords = [Set.new(), Set.new()]
    current_x = 0
    current_y = 0
    instructions.each_with_index do |wire, j|
        wire.each_with_index do |length, i|
            if i == 0
                current_x = 0
                current_y = 0
            end
            case length[:direction]
            when "R"
                (current_x..current_x + length[:length]).each do |l|
                    all_coords[j].add([l, current_y])
                end
                current_x += length[:length]
            when "L"
                current_x.downto(current_x - length[:length]).each do |l|
                    all_coords[j].add([l, current_y])
                end
                current_x -= length[:length]
            when "U"
                (current_y..current_y + length[:length]).each do |l|
                    all_coords[j].add([current_x, l])
                end
                current_y += length[:length]
            when "D"
                current_y.downto(current_y - length[:length]).each do |l|
                    all_coords[j].add([current_x, l])
                end
                current_y -= length[:length]
            end
        end
    end
    
    pairs = (all_coords[0].intersection all_coords[1]).map do |point|
        manhattan(point)
    end
    sorted = pairs.sort
    puts "Answer to part 1:", sorted[1]
end

def part_2
    all_coords = [Hash.new(), Hash.new()]
    current_x = 0
    current_y = 0
    tot_distance = 0
    instructions.each_with_index do |wire, j|
        wire.each_with_index do |length, i|
            if i == 0
                current_x = 0
                current_y = 0
                tot_distance = 0
            end
            case length[:direction]
            when "R"
                (current_x..current_x + length[:length]).each do |l|
                    if all_coords[j][[l, current_y]] == nil
                        all_coords[j][[l, current_y]] = tot_distance
                    end
                    tot_distance += 1
                end
                current_x += length[:length]
            when "L"
                current_x.downto(current_x - length[:length]).each do |l|
                    if all_coords[j][[l, current_y]] == nil
                        all_coords[j][[l, current_y]] = tot_distance
                    end
                    tot_distance += 1
                end
                current_x -= length[:length]
            when "U"
                (current_y..current_y + length[:length]).each do |l|
                    if all_coords[j][[current_x, l]] == nil
                        all_coords[j][[current_x, l]] = tot_distance
                    end
                    tot_distance += 1
                end
                current_y += length[:length]
            when "D"
                current_y.downto(current_y - length[:length]).each do |l|
                    if all_coords[j][[current_x, l]] == nil
                        all_coords[j][[current_x, l]] = tot_distance
                    end
                    tot_distance += 1
                end
                current_y -= length[:length]
            end
            tot_distance -= 1
        end
    end
    
    coords = [Set.new(), Set.new()]
    (0..1).each do |i|
        all_coords[i].keys.each do |key|
            coords[i].add(key)
        end
    end

    steps_to_intersection = []
    (coords[0] & coords[1]).each do |coord|
        steps_to_intersection << all_coords[0][coord] + all_coords[1][coord] 
    end
    sorted = steps_to_intersection.sort
    puts "Answer to part 2:", sorted[1]
end

def manhattan(point)
    x = point[0]
    y = point[1]
    x.abs + y.abs
end

part_1
part_2