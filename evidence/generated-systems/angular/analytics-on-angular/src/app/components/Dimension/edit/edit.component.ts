import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DimensionService } from '../../../services/Dimension.service';
import { SubBaseComponent } from '../../Dimension/sub.base.component';


@Component({
    selector: 'app-edit-dimension',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDimensionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Dimension';

    dimensionForm: FormGroup;
    dimension: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DimensionService,
        private fb: FormBuilder
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

    
    updateDimension(name, typeTime, SemanticModel, Datasets, GlossaryTerms, DimensionType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDimension(name, typeTime, SemanticModel, Datasets, GlossaryTerms, DimensionType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDimension']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDimension(params['id']).subscribe(res => {
                this.dimension = res;
            });
        });
    }
}