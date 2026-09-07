
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CompetencyService } from '../../../services/Competency.service';
import { Competency } from '../../../models/Competency';

@Component({
    selector: 'app-index-competency',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCompetencyComponent implements OnInit {

    competencys: Competency[] = [];

    constructor(
        private router: Router,
        private service: CompetencyService
) {}

    ngOnInit(): void {
        this.getCompetencys();
}

    getCompetencys(): void {
        this.service.getCompetencys().subscribe((res) => {
        this.competencys = res;
    });
}

    deleteCompetency(id: any): void {
        this.service.deleteCompetency(id)
            .subscribe(() => {
                this.getCompetencys();
            });
    }
}