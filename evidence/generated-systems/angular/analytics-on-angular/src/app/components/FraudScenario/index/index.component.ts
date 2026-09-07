
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FraudScenarioService } from '../../../services/FraudScenario.service';
import { FraudScenario } from '../../../models/FraudScenario';

@Component({
    selector: 'app-index-fraudScenario',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexFraudScenarioComponent implements OnInit {

    fraudScenarios: FraudScenario[] = [];

    constructor(
        private router: Router,
        private service: FraudScenarioService
) {}

    ngOnInit(): void {
        this.getFraudScenarios();
}

    getFraudScenarios(): void {
        this.service.getFraudScenarios().subscribe((res) => {
        this.fraudScenarios = res;
    });
}

    deleteFraudScenario(id: any): void {
        this.service.deleteFraudScenario(id)
            .subscribe(() => {
                this.getFraudScenarios();
            });
    }
}