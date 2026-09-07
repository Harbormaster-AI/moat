
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditInferenceEndpointComponent } from './edit.component';
import { InferenceEndpointService } from '../../../services/InferenceEndpoint.service';

describe('EditInferenceEndpointComponent', () => {
  let component: EditInferenceEndpointComponent;
  let fixture: ComponentFixture<EditInferenceEndpointComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditInferenceEndpointComponent
      ],
      providers: [
        InferenceEndpointService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditInferenceEndpointComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});