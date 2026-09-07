
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSubscriberComponent } from './create.component';
import { SubscriberService } from '../../../services/Subscriber.service';
import { Router } from '@angular/router';

describe('CreateSubscriberComponent', () => {
  let component: CreateSubscriberComponent;
  let fixture: ComponentFixture<CreateSubscriberComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSubscriberComponent
      ],
      providers: [
        SubscriberService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSubscriberComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});