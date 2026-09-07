import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SemanticModelService } from '../../../services/SemanticModel.service';
import { SemanticModel } from '../../../models/SemanticModel';
import { SubBaseComponent } from '../../SemanticModel/sub.base.component';

@Component({
    selector: 'app-create-semanticModel',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSemanticModelComponent extends SubBaseComponent implements OnInit {

    title = 'Add SemanticModel';

    semanticModelForm: FormGroup;
    semanticModel: SemanticModel;

    constructor( http: HttpClient,
        private semanticModelService: SemanticModelService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.semanticModelForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      version: ['', Validators.required],
      grain: ['', Validators.required],
      Datasets: ['', ],
      Metrics: ['', ],
      Dimensions: ['', ],
      Measures: ['', ],
      GlossaryTerms: ['', ]
        });
    }

    
    addSemanticModel(name, version, grain, Datasets, Metrics, Dimensions, Measures, GlossaryTerms): void {
        this.semanticModelService
        .addSemanticModel(name, version, grain, Datasets, Metrics, Dimensions, Measures, GlossaryTerms)
            .subscribe(() => {
                this.router.navigate(['/indexSemanticModel']);
            });
    }

    ngOnInit(): void {
    }
}