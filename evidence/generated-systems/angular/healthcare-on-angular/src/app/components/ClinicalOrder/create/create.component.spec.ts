
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateClinicalOrderComponent } from './create.component';
import { ClinicalOrderService } from '../../../services/ClinicalOrder.service';
import { Router } from '@angular/router';

describe('CreateClinicalOrderComponent', () => {
  let component: CreateClinicalOrderComponent;
  let fixture: ComponentFixture<CreateClinicalOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateClinicalOrderComponent
      ],
      providers: [
        ClinicalOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateClinicalOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});