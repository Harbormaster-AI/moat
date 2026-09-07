
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ActivityService } from '../../../services/Activity.service';
import { Activity } from '../../../models/Activity';

@Component({
    selector: 'app-index-activity',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexActivityComponent implements OnInit {

    activitys: Activity[] = [];

    constructor(
        private router: Router,
        private service: ActivityService
) {}

    ngOnInit(): void {
        this.getActivitys();
}

    getActivitys(): void {
        this.service.getActivitys().subscribe((res) => {
        this.activitys = res;
    });
}

    deleteActivity(id: any): void {
        this.service.deleteActivity(id)
            .subscribe(() => {
                this.getActivitys();
            });
    }
}