import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AircraftOptionService } from '../../../services/AircraftOption.service';
import { SubBaseComponent } from '../../AircraftOption/sub.base.component';


@Component({
    selector: 'app-edit-aircraftOption',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAircraftOptionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AircraftOption';

    aircraftOptionForm: FormGroup;
    aircraftOption: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AircraftOptionService,
        private fb: FormBuilder
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

    
    updateAircraftOption(code, name, Variants, Packages, OptionCategory): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAircraftOption(code, name, Variants, Packages, OptionCategory, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAircraftOption']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAircraftOption(params['id']).subscribe(res => {
                this.aircraftOption = res;
            });
        });
    }
}