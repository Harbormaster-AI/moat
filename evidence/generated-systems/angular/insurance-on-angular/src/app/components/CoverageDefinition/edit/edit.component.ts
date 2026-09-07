import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CoverageDefinitionService } from '../../../services/CoverageDefinition.service';
import { SubBaseComponent } from '../../CoverageDefinition/sub.base.component';


@Component({
    selector: 'app-edit-coverageDefinition',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCoverageDefinitionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CoverageDefinition';

    coverageDefinitionForm: FormGroup;
    coverageDefinition: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CoverageDefinitionService,
        private fb: FormBuilder
) {
        super(http);
        this.coverageDefinitionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      defaultLimit: ['', Validators.required],
      defaultDeductible: ['', Validators.required],
      asMandatory: ['', Validators.required],
      Product: ['', ],
      CoverageType: ['', ]
        });
    }

    
    updateCoverageDefinition(name, defaultLimit, defaultDeductible, asMandatory, Product, CoverageType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCoverageDefinition(name, defaultLimit, defaultDeductible, asMandatory, Product, CoverageType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCoverageDefinition']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCoverageDefinition(params['id']).subscribe(res => {
                this.coverageDefinition = res;
            });
        });
    }
}