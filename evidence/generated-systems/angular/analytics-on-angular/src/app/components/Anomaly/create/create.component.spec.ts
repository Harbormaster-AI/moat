
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAnomalyComponent } from './create.component';
import { AnomalyService } from '../../../services/Anomaly.service';
import { Router } from '@angular/router';

describe('CreateAnomalyComponent', () => {
  let component: CreateAnomalyComponent;
  let fixture: ComponentFixture<CreateAnomalyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAnomalyComponent
      ],
      providers: [
        AnomalyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAnomalyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});