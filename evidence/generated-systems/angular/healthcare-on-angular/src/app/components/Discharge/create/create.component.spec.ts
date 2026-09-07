
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDischargeComponent } from './create.component';
import { DischargeService } from '../../../services/Discharge.service';
import { Router } from '@angular/router';

describe('CreateDischargeComponent', () => {
  let component: CreateDischargeComponent;
  let fixture: ComponentFixture<CreateDischargeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDischargeComponent
      ],
      providers: [
        DischargeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDischargeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});