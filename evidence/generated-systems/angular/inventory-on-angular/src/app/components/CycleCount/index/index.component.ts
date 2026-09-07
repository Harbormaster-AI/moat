
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CycleCountService } from '../../../services/CycleCount.service';
import { CycleCount } from '../../../models/CycleCount';

@Component({
    selector: 'app-index-cycleCount',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCycleCountComponent implements OnInit {

    cycleCounts: CycleCount[] = [];

    constructor(
        private router: Router,
        private service: CycleCountService
) {}

    ngOnInit(): void {
        this.getCycleCounts();
}

    getCycleCounts(): void {
        this.service.getCycleCounts().subscribe((res) => {
        this.cycleCounts = res;
    });
}

    deleteCycleCount(id: any): void {
        this.service.deleteCycleCount(id)
            .subscribe(() => {
                this.getCycleCounts();
            });
    }
}