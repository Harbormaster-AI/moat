import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AircraftFamilyService } from '../../../services/AircraftFamily.service';
import { AircraftFamily } from '../../../models/AircraftFamily';
import { SubBaseComponent } from '../../AircraftFamily/sub.base.component';

@Component({
    selector: 'app-create-aircraftFamily',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAircraftFamilyComponent extends SubBaseComponent implements OnInit {

    title = 'Add AircraftFamily';

    aircraftFamilyForm: FormGroup;
    aircraftFamily: AircraftFamily;

    constructor( http: HttpClient,
        private aircraftFamilyService: AircraftFamilyService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.aircraftFamilyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      familyCode: ['', Validators.required],
      Program: ['', ],
      AircraftModels: ['', ]
        });
    }

    
    addAircraftFamily(name, familyCode, Program, AircraftModels): void {
        this.aircraftFamilyService
        .addAircraftFamily(name, familyCode, Program, AircraftModels)
            .subscribe(() => {
                this.router.navigate(['/indexAircraftFamily']);
            });
    }

    ngOnInit(): void {
    }
}