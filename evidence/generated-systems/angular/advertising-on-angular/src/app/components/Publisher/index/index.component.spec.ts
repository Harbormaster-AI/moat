
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPublisherComponent } from './index.component';
import { PublisherService } from '../../../services/Publisher.service';

describe('IndexPublisherComponent', () => {
  let component: IndexPublisherComponent;
  let fixture: ComponentFixture<IndexPublisherComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPublisherComponent
      ],
      providers: [
        PublisherService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPublisherComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});