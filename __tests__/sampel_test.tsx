import { Sample } from "@/pages/sample";
import { render, screen } from "@testing-library/react";
import React from 'react'


jest.mock('next/head', () => {
    return {
        __esModule: true,
        default: ({ children }: { children: Array<React.ReactElement> }) => {
            return <>{children}</>
        },
    }
})

describe('first test', () => {
    test('github actions jest test', async () => {
        render(<Sample />)
        // screen.debug()
        await expect(document.title).toBe('Test')
    })
})