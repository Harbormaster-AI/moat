
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateObservationComponent } from './create.component';
import { ObservationService } from '../../../services/Observation.service';
import { Router } from '@angular/router';

describe('CreateObservationComponent', () => {
  let component: CreateObservationComponent;
  let fixture: ComponentFixture<CreateObservationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateObservationComponent
      ],
      providers: [
        ObservationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateObservationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});