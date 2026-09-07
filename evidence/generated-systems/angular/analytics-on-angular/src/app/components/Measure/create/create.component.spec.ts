
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMeasureComponent } from './create.component';
import { MeasureService } from '../../../services/Measure.service';
import { Router } from '@angular/router';

describe('CreateMeasureComponent', () => {
  let component: CreateMeasureComponent;
  let fixture: ComponentFixture<CreateMeasureComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMeasureComponent
      ],
      providers: [
        MeasureService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMeasureComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});