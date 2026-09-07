
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ScheduleExceptionService } from '../../../services/ScheduleException.service';
import { ScheduleException } from '../../../models/ScheduleException';

@Component({
    selector: 'app-index-scheduleException',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexScheduleExceptionComponent implements OnInit {

    scheduleExceptions: ScheduleException[] = [];

    constructor(
        private router: Router,
        private service: ScheduleExceptionService
) {}

    ngOnInit(): void {
        this.getScheduleExceptions();
}

    getScheduleExceptions(): void {
        this.service.getScheduleExceptions().subscribe((res) => {
        this.scheduleExceptions = res;
    });
}

    deleteScheduleException(id: any): void {
        this.service.deleteScheduleException(id)
            .subscribe(() => {
                this.getScheduleExceptions();
            });
    }
}