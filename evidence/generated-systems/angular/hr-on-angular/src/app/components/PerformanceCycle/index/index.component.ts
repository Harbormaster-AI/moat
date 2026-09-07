
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PerformanceCycleService } from '../../../services/PerformanceCycle.service';
import { PerformanceCycle } from '../../../models/PerformanceCycle';

@Component({
    selector: 'app-index-performanceCycle',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPerformanceCycleComponent implements OnInit {

    performanceCycles: PerformanceCycle[] = [];

    constructor(
        private router: Router,
        private service: PerformanceCycleService
) {}

    ngOnInit(): void {
        this.getPerformanceCycles();
}

    getPerformanceCycles(): void {
        this.service.getPerformanceCycles().subscribe((res) => {
        this.performanceCycles = res;
    });
}

    deletePerformanceCycle(id: any): void {
        this.service.deletePerformanceCycle(id)
            .subscribe(() => {
                this.getPerformanceCycles();
            });
    }
}