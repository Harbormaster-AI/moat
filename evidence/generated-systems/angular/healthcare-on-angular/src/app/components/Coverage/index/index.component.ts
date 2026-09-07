
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CoverageService } from '../../../services/Coverage.service';
import { Coverage } from '../../../models/Coverage';

@Component({
    selector: 'app-index-coverage',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCoverageComponent implements OnInit {

    coverages: Coverage[] = [];

    constructor(
        private router: Router,
        private service: CoverageService
) {}

    ngOnInit(): void {
        this.getCoverages();
}

    getCoverages(): void {
        this.service.getCoverages().subscribe((res) => {
        this.coverages = res;
    });
}

    deleteCoverage(id: any): void {
        this.service.deleteCoverage(id)
            .subscribe(() => {
                this.getCoverages();
            });
    }
}