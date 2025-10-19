# Architecture Documentation

## System Design

### High-Level Architecture
┌─────────────────────────────────────────┐
│         UI Layer (Fyne)                 │
│  ├─ Main Window                         │
│  ├─ Dashboard Tab                       │
│  ├─ Results Tab                         │
│  └─ Settings Tab                        │
└────────────────┬────────────────────────┘
│
┌────────────────▼────────────────────────┐
│      Application Logic Layer            │
│  ├─ Cache Module                        │
│  ├─ Cleaner Module                      │
│  ├─ System Module                       │
│  └─ Models Module                       │
└────────────────┬────────────────────────┘
│
┌────────────────▼────────────────────────┐
│      System Interface Layer             │
│  ├─ File System Access                  │
│  ├─ Disk Usage Information              │
│  └─ Platform Detection                  │
└─────────────────────────────────────────┘

### Component Responsibilities

**UI Layer:**
- User interface rendering
- Event handling
- Progress visualization

**Application Logic:**
- Cache detection
- File operations
- Data processing

**System Interface:**
- OS-specific operations
- File I/O
- Permission handling

## Data Flow

1. User initiates scan
2. Scanner discovers cache locations
3. Results displayed in UI
4. User selects items to clean
5. Cleaner processes deletions
6. Progress updated in real-time
7. Summary report generated

## Performance Considerations

- Concurrent file scanning
- Efficient memory usage
- Progress streaming
- Cancellation support
- Error recovery