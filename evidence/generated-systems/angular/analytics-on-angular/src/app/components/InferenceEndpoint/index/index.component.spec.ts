
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInferenceEndpointComponent } from './index.component';
import { InferenceEndpointService } from '../../../services/InferenceEndpoint.service';

describe('IndexInferenceEndpointComponent', () => {
  let component: IndexInferenceEndpointComponent;
  let fixture: ComponentFixture<IndexInferenceEndpointComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInferenceEndpointComponent
      ],
      providers: [
        InferenceEndpointService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInferenceEndpointComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});