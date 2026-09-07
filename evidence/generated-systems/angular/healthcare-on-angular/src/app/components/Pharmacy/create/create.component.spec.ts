
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePharmacyComponent } from './create.component';
import { PharmacyService } from '../../../services/Pharmacy.service';
import { Router } from '@angular/router';

describe('CreatePharmacyComponent', () => {
  let component: CreatePharmacyComponent;
  let fixture: ComponentFixture<CreatePharmacyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePharmacyComponent
      ],
      providers: [
        PharmacyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePharmacyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});