
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DemandSignalService } from '../../../services/DemandSignal.service';
import { DemandSignal } from '../../../models/DemandSignal';

@Component({
    selector: 'app-index-demandSignal',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDemandSignalComponent implements OnInit {

    demandSignals: DemandSignal[] = [];

    constructor(
        private router: Router,
        private service: DemandSignalService
) {}

    ngOnInit(): void {
        this.getDemandSignals();
}

    getDemandSignals(): void {
        this.service.getDemandSignals().subscribe((res) => {
        this.demandSignals = res;
    });
}

    deleteDemandSignal(id: any): void {
        this.service.deleteDemandSignal(id)
            .subscribe(() => {
                this.getDemandSignals();
            });
    }
}