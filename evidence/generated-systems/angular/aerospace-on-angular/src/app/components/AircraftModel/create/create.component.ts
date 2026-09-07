import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AircraftModelService } from '../../../services/AircraftModel.service';
import { AircraftModel } from '../../../models/AircraftModel';
import { SubBaseComponent } from '../../AircraftModel/sub.base.component';

@Component({
    selector: 'app-create-aircraftModel',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAircraftModelComponent extends SubBaseComponent implements OnInit {

    title = 'Add AircraftModel';

    aircraftModelForm: FormGroup;
    aircraftModel: AircraftModel;

    constructor( http: HttpClient,
        private aircraftModelService: AircraftModelService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.aircraftModelForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      modelDesignation: ['', Validators.required],
      Family: ['', ],
      Variants: ['', ],
      EngineTypes: ['', ],
      AircraftType: ['', ]
        });
    }

    
    addAircraftModel(name, modelDesignation, Family, Variants, EngineTypes, AircraftType): void {
        this.aircraftModelService
        .addAircraftModel(name, modelDesignation, Family, Variants, EngineTypes, AircraftType)
            .subscribe(() => {
                this.router.navigate(['/indexAircraftModel']);
            });
    }

    ngOnInit(): void {
    }
}