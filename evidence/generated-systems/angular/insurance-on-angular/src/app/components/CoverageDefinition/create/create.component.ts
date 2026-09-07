import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CoverageDefinitionService } from '../../../services/CoverageDefinition.service';
import { CoverageDefinition } from '../../../models/CoverageDefinition';
import { SubBaseComponent } from '../../CoverageDefinition/sub.base.component';

@Component({
    selector: 'app-create-coverageDefinition',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCoverageDefinitionComponent extends SubBaseComponent implements OnInit {

    title = 'Add CoverageDefinition';

    coverageDefinitionForm: FormGroup;
    coverageDefinition: CoverageDefinition;

    constructor( http: HttpClient,
        private coverageDefinitionService: CoverageDefinitionService,
        private fb: FormBuilder,
        private router: Router
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

    
    addCoverageDefinition(name, defaultLimit, defaultDeductible, asMandatory, Product, CoverageType): void {
        this.coverageDefinitionService
        .addCoverageDefinition(name, defaultLimit, defaultDeductible, asMandatory, Product, CoverageType)
            .subscribe(() => {
                this.router.navigate(['/indexCoverageDefinition']);
            });
    }

    ngOnInit(): void {
    }
}