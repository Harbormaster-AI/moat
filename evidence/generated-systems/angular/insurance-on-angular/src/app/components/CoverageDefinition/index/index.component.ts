
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CoverageDefinitionService } from '../../../services/CoverageDefinition.service';
import { CoverageDefinition } from '../../../models/CoverageDefinition';

@Component({
    selector: 'app-index-coverageDefinition',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCoverageDefinitionComponent implements OnInit {

    coverageDefinitions: CoverageDefinition[] = [];

    constructor(
        private router: Router,
        private service: CoverageDefinitionService
) {}

    ngOnInit(): void {
        this.getCoverageDefinitions();
}

    getCoverageDefinitions(): void {
        this.service.getCoverageDefinitions().subscribe((res) => {
        this.coverageDefinitions = res;
    });
}

    deleteCoverageDefinition(id: any): void {
        this.service.deleteCoverageDefinition(id)
            .subscribe(() => {
                this.getCoverageDefinitions();
            });
    }
}