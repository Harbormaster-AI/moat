
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSubscriberComponent } from './index.component';
import { SubscriberService } from '../../../services/Subscriber.service';

describe('IndexSubscriberComponent', () => {
  let component: IndexSubscriberComponent;
  let fixture: ComponentFixture<IndexSubscriberComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSubscriberComponent
      ],
      providers: [
        SubscriberService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSubscriberComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});