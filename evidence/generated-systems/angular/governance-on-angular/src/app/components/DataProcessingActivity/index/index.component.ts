
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DataProcessingActivityService } from '../../../services/DataProcessingActivity.service';
import { DataProcessingActivity } from '../../../models/DataProcessingActivity';

@Component({
    selector: 'app-index-dataProcessingActivity',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDataProcessingActivityComponent implements OnInit {

    dataProcessingActivitys: DataProcessingActivity[] = [];

    constructor(
        private router: Router,
        private service: DataProcessingActivityService
) {}

    ngOnInit(): void {
        this.getDataProcessingActivitys();
}

    getDataProcessingActivitys(): void {
        this.service.getDataProcessingActivitys().subscribe((res) => {
        this.dataProcessingActivitys = res;
    });
}

    deleteDataProcessingActivity(id: any): void {
        this.service.deleteDataProcessingActivity(id)
            .subscribe(() => {
                this.getDataProcessingActivitys();
            });
    }
}