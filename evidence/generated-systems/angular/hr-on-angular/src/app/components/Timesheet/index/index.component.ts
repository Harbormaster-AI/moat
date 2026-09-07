
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TimesheetService } from '../../../services/Timesheet.service';
import { Timesheet } from '../../../models/Timesheet';

@Component({
    selector: 'app-index-timesheet',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTimesheetComponent implements OnInit {

    timesheets: Timesheet[] = [];

    constructor(
        private router: Router,
        private service: TimesheetService
) {}

    ngOnInit(): void {
        this.getTimesheets();
}

    getTimesheets(): void {
        this.service.getTimesheets().subscribe((res) => {
        this.timesheets = res;
    });
}

    deleteTimesheet(id: any): void {
        this.service.deleteTimesheet(id)
            .subscribe(() => {
                this.getTimesheets();
            });
    }
}