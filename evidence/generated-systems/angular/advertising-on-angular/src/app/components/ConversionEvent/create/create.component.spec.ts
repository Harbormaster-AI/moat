
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateConversionEventComponent } from './create.component';
import { ConversionEventService } from '../../../services/ConversionEvent.service';
import { Router } from '@angular/router';

describe('CreateConversionEventComponent', () => {
  let component: CreateConversionEventComponent;
  let fixture: ComponentFixture<CreateConversionEventComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateConversionEventComponent
      ],
      providers: [
        ConversionEventService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateConversionEventComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});