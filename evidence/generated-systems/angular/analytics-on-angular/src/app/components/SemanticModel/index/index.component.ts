
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SemanticModelService } from '../../../services/SemanticModel.service';
import { SemanticModel } from '../../../models/SemanticModel';

@Component({
    selector: 'app-index-semanticModel',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSemanticModelComponent implements OnInit {

    semanticModels: SemanticModel[] = [];

    constructor(
        private router: Router,
        private service: SemanticModelService
) {}

    ngOnInit(): void {
        this.getSemanticModels();
}

    getSemanticModels(): void {
        this.service.getSemanticModels().subscribe((res) => {
        this.semanticModels = res;
    });
}

    deleteSemanticModel(id: any): void {
        this.service.deleteSemanticModel(id)
            .subscribe(() => {
                this.getSemanticModels();
            });
    }
}