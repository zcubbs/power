import {CaretSortIcon, CheckIcon} from "@radix-ui/react-icons"

import {cn} from "@/lib/utils"
import {Button} from "@/components/ui/button"
import {Command, CommandEmpty, CommandGroup, CommandInput, CommandItem,} from "@/components/ui/command"
import {Popover, PopoverContent, PopoverTrigger,} from "@/components/ui/popover"
import React from "react";

interface ComboboxProps {
  defaultValue: string
  placeholder: string
  onChange: (value: string) => void
  options: string[]
}

export const Combobox: React.FC<ComboboxProps> = ({
                                                    defaultValue,
                                                    placeholder,
                                                    onChange,
                                                    options,
                                                  }) => {
  const [open, setOpen] = React.useState(false)
  const [selectedValue, setSelectedValue] = React.useState(defaultValue)

  // set default value
  React.useEffect(() => {
    if (defaultValue) {
      setSelectedValue(defaultValue)
    }
  }, [defaultValue])

  const onSelectedValueChange =(value: string) => {
    if (value === selectedValue) return
    setSelectedValue(value)
    setOpen(false)
    onChange(value)
  }

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="ul"
          aria-expanded={open}
          className="w-[200px] justify-between"
        >
          {selectedValue
            ? options?.find((option) => option === selectedValue)
            : placeholder}
          <CaretSortIcon className="ml-2 h-4 w-4 shrink-0 opacity-50"/>
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[200px] p-0">
        <Command>
          <CommandInput placeholder={placeholder} className="h-9"/>
          <CommandEmpty>No options.</CommandEmpty>
          <CommandGroup>
            {options?.map((option) => (
              <CommandItem
                key={option}
                value={option}
                onSelect={(_) => {
                  onSelectedValueChange(option);
                }}
              >
                {option}
                <CheckIcon
                  className={cn(
                    "ml-auto h-4 w-4",
                    selectedValue === option ? "opacity-100" : "opacity-0"
                  )}
                />
              </CommandItem>
            ))}
          </CommandGroup>
        </Command>
      </PopoverContent>
    </Popover>
  )
}
