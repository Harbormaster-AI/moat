
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAPIClientComponent } from './index.component';
import { APIClientService } from '../../../services/APIClient.service';

describe('IndexAPIClientComponent', () => {
  let component: IndexAPIClientComponent;
  let fixture: ComponentFixture<IndexAPIClientComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAPIClientComponent
      ],
      providers: [
        APIClientService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAPIClientComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});