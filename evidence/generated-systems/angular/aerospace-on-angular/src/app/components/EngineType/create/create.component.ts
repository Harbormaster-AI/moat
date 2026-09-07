import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EngineTypeService } from '../../../services/EngineType.service';
import { EngineType } from '../../../models/EngineType';
import { SubBaseComponent } from '../../EngineType/sub.base.component';

@Component({
    selector: 'app-create-engineType',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateEngineTypeComponent extends SubBaseComponent implements OnInit {

    title = 'Add EngineType';

    engineTypeForm: FormGroup;
    engineType: EngineType;

    constructor( http: HttpClient,
        private engineTypeService: EngineTypeService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.engineTypeForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  engineModelCode: ['', Validators.required],
      maxThrustKn: ['', Validators.required],
      Supplier: ['', ],
      CompatibleModels: ['', ],
      Category: ['', ]
        });
    }

    
    addEngineType(engineModelCode, maxThrustKn, Supplier, CompatibleModels, Category): void {
        this.engineTypeService
        .addEngineType(engineModelCode, maxThrustKn, Supplier, CompatibleModels, Category)
            .subscribe(() => {
                this.router.navigate(['/indexEngineType']);
            });
    }

    ngOnInit(): void {
    }
}