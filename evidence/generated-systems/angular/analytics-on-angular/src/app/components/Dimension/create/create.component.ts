import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DimensionService } from '../../../services/Dimension.service';
import { Dimension } from '../../../models/Dimension';
import { SubBaseComponent } from '../../Dimension/sub.base.component';

@Component({
    selector: 'app-create-dimension',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDimensionComponent extends SubBaseComponent implements OnInit {

    title = 'Add Dimension';

    dimensionForm: FormGroup;
    dimension: Dimension;

    constructor( http: HttpClient,
        private dimensionService: DimensionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.dimensionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      typeTime: ['', Validators.required],
      SemanticModel: ['', ],
      Datasets: ['', ],
      GlossaryTerms: ['', ],
      DimensionType: ['', ]
        });
    }

    
    addDimension(name, typeTime, SemanticModel, Datasets, GlossaryTerms, DimensionType): void {
        this.dimensionService
        .addDimension(name, typeTime, SemanticModel, Datasets, GlossaryTerms, DimensionType)
            .subscribe(() => {
                this.router.navigate(['/indexDimension']);
            });
    }

    ngOnInit(): void {
    }
}