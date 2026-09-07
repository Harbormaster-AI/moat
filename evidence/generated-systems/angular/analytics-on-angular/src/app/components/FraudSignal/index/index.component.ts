
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FraudSignalService } from '../../../services/FraudSignal.service';
import { FraudSignal } from '../../../models/FraudSignal';

@Component({
    selector: 'app-index-fraudSignal',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexFraudSignalComponent implements OnInit {

    fraudSignals: FraudSignal[] = [];

    constructor(
        private router: Router,
        private service: FraudSignalService
) {}

    ngOnInit(): void {
        this.getFraudSignals();
}

    getFraudSignals(): void {
        this.service.getFraudSignals().subscribe((res) => {
        this.fraudSignals = res;
    });
}

    deleteFraudSignal(id: any): void {
        this.service.deleteFraudSignal(id)
            .subscribe(() => {
                this.getFraudSignals();
            });
    }
}