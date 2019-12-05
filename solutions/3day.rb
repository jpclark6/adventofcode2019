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

def manhattan(point)
    x = point[0]
    y = point[1]
    x.abs + y.abs
end

part_1