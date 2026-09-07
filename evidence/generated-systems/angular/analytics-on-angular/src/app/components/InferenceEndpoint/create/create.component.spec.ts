
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInferenceEndpointComponent } from './create.component';
import { InferenceEndpointService } from '../../../services/InferenceEndpoint.service';
import { Router } from '@angular/router';

describe('CreateInferenceEndpointComponent', () => {
  let component: CreateInferenceEndpointComponent;
  let fixture: ComponentFixture<CreateInferenceEndpointComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInferenceEndpointComponent
      ],
      providers: [
        InferenceEndpointService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInferenceEndpointComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});