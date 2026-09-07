
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInsuranceProductComponent } from './create.component';
import { InsuranceProductService } from '../../../services/InsuranceProduct.service';
import { Router } from '@angular/router';

describe('CreateInsuranceProductComponent', () => {
  let component: CreateInsuranceProductComponent;
  let fixture: ComponentFixture<CreateInsuranceProductComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInsuranceProductComponent
      ],
      providers: [
        InsuranceProductService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInsuranceProductComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});