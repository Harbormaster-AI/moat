
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RegulationService } from '../../../services/Regulation.service';
import { Regulation } from '../../../models/Regulation';

@Component({
    selector: 'app-index-regulation',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRegulationComponent implements OnInit {

    regulations: Regulation[] = [];

    constructor(
        private router: Router,
        private service: RegulationService
) {}

    ngOnInit(): void {
        this.getRegulations();
}

    getRegulations(): void {
        this.service.getRegulations().subscribe((res) => {
        this.regulations = res;
    });
}

    deleteRegulation(id: any): void {
        this.service.deleteRegulation(id)
            .subscribe(() => {
                this.getRegulations();
            });
    }
}