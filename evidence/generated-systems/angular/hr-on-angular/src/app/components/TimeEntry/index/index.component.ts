
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TimeEntryService } from '../../../services/TimeEntry.service';
import { TimeEntry } from '../../../models/TimeEntry';

@Component({
    selector: 'app-index-timeEntry',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTimeEntryComponent implements OnInit {

    timeEntrys: TimeEntry[] = [];

    constructor(
        private router: Router,
        private service: TimeEntryService
) {}

    ngOnInit(): void {
        this.getTimeEntrys();
}

    getTimeEntrys(): void {
        this.service.getTimeEntrys().subscribe((res) => {
        this.timeEntrys = res;
    });
}

    deleteTimeEntry(id: any): void {
        this.service.deleteTimeEntry(id)
            .subscribe(() => {
                this.getTimeEntrys();
            });
    }
}