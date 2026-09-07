
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateExperimentVariantComponent } from './create.component';
import { ExperimentVariantService } from '../../../services/ExperimentVariant.service';
import { Router } from '@angular/router';

describe('CreateExperimentVariantComponent', () => {
  let component: CreateExperimentVariantComponent;
  let fixture: ComponentFixture<CreateExperimentVariantComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateExperimentVariantComponent
      ],
      providers: [
        ExperimentVariantService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateExperimentVariantComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});