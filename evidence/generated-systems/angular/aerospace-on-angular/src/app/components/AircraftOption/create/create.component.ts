import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AircraftOptionService } from '../../../services/AircraftOption.service';
import { AircraftOption } from '../../../models/AircraftOption';
import { SubBaseComponent } from '../../AircraftOption/sub.base.component';

@Component({
    selector: 'app-create-aircraftOption',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAircraftOptionComponent extends SubBaseComponent implements OnInit {

    title = 'Add AircraftOption';

    aircraftOptionForm: FormGroup;
    aircraftOption: AircraftOption;

    constructor( http: HttpClient,
        private aircraftOptionService: AircraftOptionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.aircraftOptionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  code: ['', Validators.required],
      name: ['', Validators.required],
      Variants: ['', ],
      Packages: ['', ],
      OptionCategory: ['', ]
        });
    }

    
    addAircraftOption(code, name, Variants, Packages, OptionCategory): void {
        this.aircraftOptionService
        .addAircraftOption(code, name, Variants, Packages, OptionCategory)
            .subscribe(() => {
                this.router.navigate(['/indexAircraftOption']);
            });
    }

    ngOnInit(): void {
    }
}