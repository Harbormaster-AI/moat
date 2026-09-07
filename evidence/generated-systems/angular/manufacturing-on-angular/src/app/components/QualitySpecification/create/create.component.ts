import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { QualitySpecificationService } from '../../../services/QualitySpecification.service';
import { QualitySpecification } from '../../../models/QualitySpecification';
import { SubBaseComponent } from '../../QualitySpecification/sub.base.component';

@Component({
    selector: 'app-create-qualitySpecification',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateQualitySpecificationComponent extends SubBaseComponent implements OnInit {

    title = 'Add QualitySpecification';

    qualitySpecificationForm: FormGroup;
    qualitySpecification: QualitySpecification;

    constructor( http: HttpClient,
        private qualitySpecificationService: QualitySpecificationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.qualitySpecificationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  specCode: ['', Validators.required],
      name: ['', Validators.required],
      version: ['', Validators.required],
      Item: ['', ]
        });
    }

    
    addQualitySpecification(specCode, name, version, Item): void {
        this.qualitySpecificationService
        .addQualitySpecification(specCode, name, version, Item)
            .subscribe(() => {
                this.router.navigate(['/indexQualitySpecification']);
            });
    }

    ngOnInit(): void {
    }
}